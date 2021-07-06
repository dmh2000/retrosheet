const { default: axios } = require("axios");

function resolveUser(parentValue, args) {
  return axios.get(`http://localhost:3000/user/${args.id}`).then((rsp) => {
    return rsp.data;
  });
}

function resolveCompany(parentValue, args) {
  console.log(parentValue, args);
  return axios.get(`http://localhost:3000/company/${args.id}`).then((rsp) => {
    return rsp.data;
  });
}

module.exports = {
  resolveUser,
  resolveCompany,
};
