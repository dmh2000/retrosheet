const { default: axios } = require("axios");

function resolveUser(parentValue, args) {
  return axios.get(`http://localhost:3000/user/${args.id}`).then((rsp) => {
    return rsp.data;
  });
}

function resolveUsers(parentValue, args) {
  return axios
    .get(`http://localhost:3000/company/${args.id}/user`)
    .then((rsp) => {
      return rsp.data;
    });
}

function resolveCompany(parentValue, args) {
  return axios.get(`http://localhost:3000/company/${args.id}`).then((rsp) => {
    return rsp.data;
  });
}

function resolveAddUser(parentValue, { name, age }) {
  return axios
    .post(`http://localhost:3000/user`, {
      name,
      age,
    })
    .then((rsp) => {
      return rsp.data;
    });
}

function resolveDeleteUser(parentValue, { id }) {
  return axios.delete(`http://localhost:3000/user/${id}`).then((rsp) => {
    return rsp.data;
  });
}

function resolveModifyUser(parentValue, { id, name, age }) {
  const args = {};
  if (name) {
    args.name = name;
  }
  if (age) {
    args.age = age;
  }
  console.log(id, name, age);
  return axios.patch(`http://localhost:3000/user/${id}`, args).then((rsp) => {
    return rsp.data;
  });
}

module.exports = {
  resolveUser,
  resolveCompany,
  resolveUsers,
  resolveAddUser,
  resolveDeleteUser,
  resolveModifyUser,
};
