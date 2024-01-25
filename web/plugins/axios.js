export default function ({ $axios, $config }) {
  // timeout for request
  $axios.defaults.timeout = 1000 * 59; // 59 seconds

  if (process.server) {
    $axios.setBaseURL($config.baseURL);
  }

  // Exluding 2xx
  $axios.onError((error) => {
    if (process.client && error.response.status === 401) {
      window.location.reload();
    }
  });
}
