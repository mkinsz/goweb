<!DOCTYPE html>
<html lang="zh">

<head>
    <style>
        body {
            height: 100%;
            margin: 0;
            width: 100%;
            overflow: hidden;
            font-family: "Segoe UI", "Roboto", "Noto Sans", "Ubuntu", "Droid Sans",
                "Helvetica Neue", sans-serif;
            font-size: 14px;
        }

        #graphiql {
            height: 100vh;
        }
    </style>
    <title>{{ .title }}</title>
    <!--
      This GraphiQL example depends on Promise and fetch, which are available in
      modern browsers, but can be "polyfilled" for older browsers.
      GraphiQL itself depends on React DOM.
      If you do not want to rely on a CDN, you can host these files locally or
      include them directly in your favored resource bunder.
    -->
    <script src="https://cdn.jsdelivr.net/npm/es6-promise@4.2.8/dist/es6-promise.auto.min.js"></script>
    <script src="https://cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/react@15.6.2/dist/react.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/react-dom@15.6.2/dist/react-dom.min.js"></script>

    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphiql@0.17.5/graphiql.min.css" />
    <script src="https://cdn.jsdelivr.net/npm/graphiql@0.17.5/graphiql.min.js"></script>
</head>

<body>
    <div id="graphiql">Loading graphiql editor...</div>
    <script>
        /**
         * This GraphiQL example illustrates how to use some of GraphiQL's props
         * in order to enable reading and updating the URL parameters, making
         * link sharing of queries a little bit easier.
         *
         * This is only one example of this kind of feature, GraphiQL exposes
         * various React params to enable interesting integrations.
         */
        // Parse the search string to get url parameters.
        var search = window.location.search;
        var parameters = {};
        search
            .substr(1)
            .split("&")
            .forEach(function (entry) {
                var eq = entry.indexOf("=");
                if (eq >= 0) {
                    parameters[
                        decodeURIComponent(entry.slice(0, eq))
                    ] = decodeURIComponent(entry.slice(eq + 1));
                }
            });
        // if variables was provided, try to format it.
        if (parameters.variables) {
            try {
                parameters.variables = JSON.stringify(
                    JSON.parse(parameters.variables),
                    null,
                    2
                );
            } catch (e) {
                // Do nothing, we want to display the invalid JSON as a string, rather
                // than present an error.
            }
        }
        // When the query and variables string is edited, update the URL bar so
        // that it can be easily shared
        function onEditQuery(newQuery) {
            parameters.query = newQuery;
            updateURL();
        }

        function onEditVariables(newVariables) {
            parameters.variables = newVariables;
            updateURL();
        }

        function onEditOperationName(newOperationName) {
            parameters.operationName = newOperationName;
            updateURL();
        }

        function updateURL() {
            var newSearch =
                "?" +
                Object.keys(parameters)
                    .filter(function (key) {
                        return Boolean(parameters[key]);
                    })
                    .map(function (key) {
                        return (
                            encodeURIComponent(key) +
                            "=" +
                            encodeURIComponent(parameters[key])
                        );
                    })
                    .join("&");
            history.replaceState(null, null, newSearch);
        }

        // Defines a GraphQL fetcher using the fetch API. You're not required to
        // use fetch, and could instead implement graphQLFetcher however you like,
        // as long as it returns a Promise or Observable.
        function graphQLFetcher(graphQLParams) {
            // This example expects a GraphQL server at the path /graphql.
            // Change this to point wherever you host your GraphQL server.
			console.log("Graphql Fetch: ", graphQLParams)
            return fetch("/graphql", {
                method: "post",
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(graphQLParams),
 
            })
                .then(function (response) {
                    return response.text();
                })
                .then(function (responseBody) {
                    try {
                        return JSON.parse(responseBody);
                    } catch (error) {
                        return responseBody;
                    }
                });
        }

        // Render <GraphiQL /> into the body.
        // See the README in the top level of this module to learn more about
        // how you can customize GraphiQL by providing different values or
        // additional child elements.
        ReactDOM.render(
            React.createElement(GraphiQL, {
                fetcher: graphQLFetcher,
                query: parameters.query,
                variables: parameters.variables,
                operationName: parameters.operationName,
                onEditQuery: onEditQuery,
                onEditVariables: onEditVariables,
                onEditOperationName: onEditOperationName,
                editorTheme: "solarized"
            }),
            document.getElementById("graphiql")
        );
    </script>
</body>

</html>