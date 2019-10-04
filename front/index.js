const { ApolloServer, gql } = require('apollo-server');

const typeDefs = gql`
    # Comment
    
    type Book {
        title: String
        author: String
    }
    
    type Query {
        books: [Book]
    }    
`;

const books = [
    {
        title: 'たいとる１つ目',
        author: 'Sato',
    },
    {
        title: 'たいとる２つ目',
        author: 'Watanabe',
    }
];

const resolvers = {
    Query: {
        books: () => books,
    },
};

const server = new ApolloServer({ typeDefs, resolvers });

server.listen().then(({ url }) => {
    console.log(`🚀Server ready at ${url}`);
});
