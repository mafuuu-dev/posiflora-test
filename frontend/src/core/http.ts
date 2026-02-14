import Singleton from "@core/singleton";
import axios, { type AxiosError, type AxiosInstance, type AxiosResponse } from "axios";

class ResponseHandler extends Singleton {
  handle = (response: AxiosResponse): AxiosResponse => response;
}

class ErrorHandler extends Singleton {
  handle = (error: AxiosError) => Promise.reject(error);
}

class HTTP extends Singleton {
  axios(): AxiosInstance {
    const instance = axios.create({
      baseURL: `${process.env.NEXT_PUBLIC_URL}/api`,
      headers: this.getHeaders(),
    });

    return this.withInterceptors(instance);
  }

  private getHeaders(): object {
    return {
      "Accept": "application/json",
      "Content-Type": "application/json",
    };
  }

  private withInterceptors(instance: AxiosInstance): AxiosInstance {
    instance.interceptors.response.use(
      (response: AxiosResponse) => ResponseHandler.getInstance().handle(response),
      (error: AxiosError) => ErrorHandler.getInstance().handle(error),
    );

    return instance;
  }
}

export default HTTP;
