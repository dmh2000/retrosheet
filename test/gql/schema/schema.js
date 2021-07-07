const {
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString,
  GraphQLInt,
  GraphQLNonNull,
} = require("graphql");
const {
  resolveUser,
  resolveCompany,
  resolveAddUser,
  resolveDeleteUser,
  resolveModifyUser,
} = require("./resolvers");

const { UserType, CompanyType } = require("./types");

const Query = new GraphQLObjectType({
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

const Mutation = new GraphQLObjectType({
  name: "Mutation",
  fields: {
    addUser: {
      type: UserType,
      args: {
        name: { type: GraphQLNonNull(GraphQLString) },
        age: { type: GraphQLNonNull(GraphQLInt) },
        companyId: { type: GraphQLString },
      },
      resolve: resolveAddUser,
    },
    deleteUser: {
      type: UserType,
      args: {
        id: { type: GraphQLNonNull(GraphQLString) },
      },
      resolve: resolveDeleteUser,
    },
    modifyUser: {
      type: UserType,
      args: {
        id: { type: GraphQLNonNull(GraphQLString) },
        name: { type: GraphQLString },
        age: { type: GraphQLInt },
      },
      resolve: resolveModifyUser,
    },
  },
});

const schema = new GraphQLSchema({
  query: Query,
  mutation: Mutation,
});

module.exports = schema;
