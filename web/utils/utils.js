const utils = {
  getDaysToMilliSecond: (days) => {
    return days * 24 * 3600 * 1000
  },
  getColor: (str) => {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
      hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return `hsl(${hash % 360}, 100%, 80%)`;
  },
}

export default utils;
