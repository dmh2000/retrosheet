const {
  GraphQLObjectType,
  GraphQLString,
  GraphQLInt,
  GraphQLList,
} = require("graphql");
const { resolveCompany, resolveUsers } = require("./resolvers");

const UserType = new GraphQLObjectType({
  name: "User",
  // wrap with anonymous function to resolve circular references
  fields: () => ({
    id: {
      type: GraphQLString,
    },
    name: {
      type: GraphQLString,
    },
    age: {
      type: GraphQLInt,
    },
    company: {
      type: CompanyType,
      args: { companyId: { type: GraphQLString } },
      resolve: (parentValue, args) => {
        return resolveCompany(undefined, { id: parentValue.companyId });
      },
    },
  }),
});

const CompanyType = new GraphQLObjectType({
  name: "Company",
  // wrap with anonymous function to resolve circular references
  fields: () => ({
    id: {
      type: GraphQLString,
    },
    name: {
      type: GraphQLString,
    },
    description: {
      type: GraphQLString,
    },
    users: {
      type: new GraphQLList(UserType),
      args: undefined,
      resolve: (parentValue, args) => {
        return resolveUsers(undefined, { id: parentValue.id });
      },
    },
  }),
});

module.exports = {
  UserType,
  CompanyType,
};
