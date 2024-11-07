package com.pucrs.api.dtos;

public class ResponseDto<T> {
  private boolean success;
  private T message;

  public ResponseDto(T message) {
    this.success = true;
    this.message = message;
  }

  public ResponseDto(String message) {
    this.success = false;
    this.message = (T) message;
  }

  public ResponseDto(boolean success, T message) {
    this.success = success;
    this.message = (T) message;
  }

  public boolean isSuccess() {
    return success;
  }

  public T getMessage() {
    return message;
  }

  public void setMessage(T message) {
    this.message = message;
  }
}
