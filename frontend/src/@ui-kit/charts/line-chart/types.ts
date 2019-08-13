export interface LineData<T, U> {
  name: string;
  data: Array<CoordData<T, U>>;
}

export interface CoordData<T, U> {
  x: T;
  y: U;
}
