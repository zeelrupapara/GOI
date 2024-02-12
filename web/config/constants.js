const CONSTANTS = {
  API_URL_PREFIX: "/api/v1",
  PRODUCT_NAME: "GPAT",
  GITHUB_URL_PREFIX: "https://github.com",

  // Filters
  FILTERS: {

    // placeholders
    ORG_PLACEHOLDER: "Select Organizations",
    REPO_PLACEHOLDER: "Select Repositories",
    MEMBER_PLACEHOLDER: "Select Members",
    DATETIME_PLACEOLDER: "Select Date Range",

    // queryparams (qp) filter keys
    ORG_QP: "orgs",
    REPO_QP: "repo",
    MEMBER_QP: "membs",
  },

  // charts
  CHART:{
    ORG_CHART_TITLE: "Organizations Contribution",

    // colors
    OPEN_PR_COLOR: "rgb(34, 134, 58, 0.6)",
    CLOSED_PR_COLOR: "rgb(214, 23, 38, 0.6)",
    MERGED_PR_COLOR: "rgb(136, 23, 152, 0.6)",
    OPEN_ISSUE_COLOR: "rgb(34, 134, 58, 0.6)",
    CLOSED_ISSUE_COLOR: "rgb(214, 23, 38, 0.6)",
  },

  // Messages
  MESSAGES: {
    GITHUB_DATA_FETCH_ERROR: "Failed to fetch data from GitHub. Please try again later.",
    GITHUB_DATA_FETCH_SUCCESS: "Process to fetch data from GitHub, please wait. This may take a few minutes.",
  },

  STATUS_CODES: {
    100: "Continue",
    101: "Switching Protocols",
    102: "Processing",
    103: "Early Hints",
    200: "OK",
    201: "Created",
    202: "Accepted",
    204: "No Content",
    205: "Reset Content",
    206: "Partial Content",
    207: "Multi-Status",
    208: "Already Reported",
    226: "IM Used",
    300: "Multiple Choices",
    301: "Moved Permanently",
    302: "Found",
    303: "See Other",
    304: "Not Modified",
    305: "Use Proxy",
    306: "Switch Proxy",
    307: "Temporary Redirect",
    308: "Permanent Redirect",
    400: "Bad Request",
    401: "Unauthorized",
    402: "Payment Required",
    403: "Forbidden",
    404: "Not Found",
    405: "Method Not Allowed",
    406: "Not Acceptable",
    407: "Proxy Authentication Required",
    408: "Request Timeout",
    409: "Conflict",
    410: "Gone",
    411: "Length Required",
    412: "Precondition Failed",
    413: "Payload Too Large",
    414: "URI Too Long",
    415: "Unsupported Media Type",
    416: "Range Not Satisfiable",
    417: "Expectation Failed",
    418: "I'm a teapot",
    421: "Misdirected Request",
    422: "Unprocessable Entity",
    423: "Locked",
    424: "Failed Dependency",
    425: "Unordered Collection",
    426: "Upgrade Required",
    428: "Precondition Required",
    429: "Too Many Requests",
    431: "Request Header Fields Too Large",
    451: "Unavailable For Legal Reasons",
    500: "Internal Server Error",
    502: "Bad Gateway",
    503: "Service Unavailable",
    504: "Gateway Timeout",
    505: "HTTP Version Not Supported",
    506: "Variant Also Negotiates",
    507: "Insufficient Storage",
    508: "Loop Detected",
    510: "Not Extended",
    511: "Network Authentication Required"
  },
};
export default CONSTANTS;
