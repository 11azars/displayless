import { withHooks, useState } from 'vue-hooks';
import gql from 'graphql-tag';

import client from './graphql';
import './app.css';

/* eslint-disable no-unused-vars */
export default withHooks((h) => {
  const [networks, setNetworks] = useState([]);
  client.query({
    query: gql`
      query {
        GetWirelessNetworks {
          SSID
        }
      }
    `,
  })
    .then(data => console.log)
    .catch(error => console.error(error));

  return (
    <div id="app">
      Hello World { networks }
    </div>
  );
});
