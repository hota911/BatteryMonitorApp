/*
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */
/*
 * This code was generated by https://github.com/google/apis-client-generator/
 * (build: 2016-01-08 17:48:37 UTC)
 * on 2016-02-01 at 15:41:07 UTC 
 * Modify at your own risk.
 */

package com.appspot.icumn7abiu.batteryservice.model;

/**
 * Model definition for HelloResp.
 *
 * <p> This is the Java data model class that specifies how to parse/serialize into the JSON that is
 * transmitted over HTTP when working with the batteryservice. For a detailed explanation see:
 * <a href="https://developers.google.com/api-client-library/java/google-http-java-client/json">https://developers.google.com/api-client-library/java/google-http-java-client/json</a>
 * </p>
 *
 * @author Google, Inc.
 */
@SuppressWarnings("javadoc")
public final class HelloResp extends com.google.api.client.json.GenericJson {

  /**
   * The value may be {@code null}.
   */
  @com.google.api.client.util.Key("Response")
  private java.lang.String response;

  /**
   * @return value or {@code null} for none
   */
  public java.lang.String getResponse() {
    return response;
  }

  /**
   * @param response response or {@code null} for none
   */
  public HelloResp setResponse(java.lang.String response) {
    this.response = response;
    return this;
  }

  @Override
  public HelloResp set(String fieldName, Object value) {
    return (HelloResp) super.set(fieldName, value);
  }

  @Override
  public HelloResp clone() {
    return (HelloResp) super.clone();
  }

}
