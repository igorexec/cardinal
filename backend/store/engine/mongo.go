package engine

import (
	"github.com/globalsign/mgo"
	"github.com/go-pkgz/mongo"
	"github.com/hashicorp/go-multierror"
	"github.com/icheliadinski/cardinal/store"
	"github.com/pkg/errors"
	"time"
)

const (
	mongoPagespeeds     = "pagespeed"
	mongoMetaPagespeeds = "pagespeed_meta"
)

type Mongo struct {
	conn            *mongo.Connection
	pagespeedWriter mongo.BufferedWriter
}

func NewMongo(conn *mongo.Connection, bufferSize int, flushDuration time.Duration) (*Mongo, error) {
	writer := mongo.NewBufferedWriter(bufferSize, conn).WithCollection(mongoPagespeeds).WithAutoFlush(flushDuration)
	result := Mongo{conn: conn, pagespeedWriter: writer}
	err := result.prepare()
	return &result, errors.Wrap(err, "failed to prepare mongo")
}

func (m *Mongo) Save(pageSpeed store.PageSpeed) (pageSpeedID string, err error) {
	err = m.conn.WithCustomCollection(mongoPagespeeds, func(coll *mgo.Collection) error {
		return coll.Insert(&pageSpeed)
	})
	return pageSpeed.Page, err
}

func (m *Mongo) prepare() error {
	errs := new(multierror.Error)
	e := m.conn.WithCustomCollection(mongoPagespeeds, func(coll *mgo.Collection) error {
		errs = multierror.Append(errs, coll.EnsureIndexKey("_id", "page", "date"))
		return errors.Wrapf(errs.ErrorOrNil(), "can't create index for %s", mongoPagespeeds)
	})
	if e != nil {
		return e
	}

	return m.conn.WithCustomCollection(mongoMetaPagespeeds, func(coll *mgo.Collection) error {
		errs = multierror.Append(errs, coll.EnsureIndexKey("_id", "page"))
		return errors.Wrapf(errs.ErrorOrNil(), "can't create index for %s", mongoMetaPagespeeds)
	})
}
