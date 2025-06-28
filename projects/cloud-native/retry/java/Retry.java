package retry.java;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class Retry {

    public void retry(RetryConfig retryConfig, ApiCall apiCall) {
        HttpClient client = HttpClient.newHttpClient();
        HttpRequest req = HttpRequest.newBuilder().uri(URI.create(apiCall.url)).build();
        client.send(req, HttpResponse.BodyHandlers.ofString());
    }


    public static class RetryConfig {
        private int maxRetries;
        private long delay;
    }

    public static class ApiCall {
        private String url;
        private String method;
    }
}
