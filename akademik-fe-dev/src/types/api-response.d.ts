interface ApiResponse<T> {
  error: boolean;
  data: T;
  message: string;
  status: number;
}
