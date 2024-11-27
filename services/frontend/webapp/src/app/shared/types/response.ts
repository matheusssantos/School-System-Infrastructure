export type Response<T = any> = {
  success: true;
  message: T;
} | {
  success: false;
  message: string;
}