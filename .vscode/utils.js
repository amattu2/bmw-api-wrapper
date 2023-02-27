const { URLSearchParams } = require('url');

/**
 * Get query param from URL
 *
 * @param {string} url positional argument URL to parse from
 * @param {string} param param to get from URL
 * @returns {string} value of param
 */
function getQueryParam(url, param) {
  const urlParams = new URLSearchParams(url);

  return urlParams.get(param) || '';
}

module.exports = [getQueryParam];
