const { default: axios } = require("axios");
const {
  GraphQLObjectType,
  GraphQLString,
  GraphQLInt,
  GraphQLSchema,
} = require("graphql");

const {
  resolveUserCompany,
  resolveUser,
  resolveCompany,
} = require("./resolvers");

const CompanyType = new GraphQLObjectType({
  name: "Company",
  fields: {
    id: {
      type: GraphQLString,
    },
    name: {
      type: GraphQLString,
    },
    description: {
      type: GraphQLString,
    },
  },
});

const UserType = new GraphQLObjectType({
  name: "User",
  fields: {
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
  },
});

const RootQuery = new GraphQLObjectType({
  name: "RootQueryType",
  fields: {
    user: {
      type: UserType,
      args: { id: { type: GraphQLString } },
      resolve: resolveUser,
    },
    company: {
      type: CompanyType,
      args: { id: { type: GraphQLString } },
      resolve: resolveCompany,
    },
  },
});

const schema = new GraphQLSchema({
  query: RootQuery,
});

module.exports = schema;
