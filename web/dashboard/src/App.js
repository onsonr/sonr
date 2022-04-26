import './App.css';
import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { Grid, Button, Message, Form, Segment, Header } from 'semantic-ui-react';
import { getProfile } from './webauthn';
import { LoginButton } from '@sonr-io/react-components';

axios.defaults.withCredentials = true;

function App() {
  const [errMsg, setErrMsg] = useState('');
  const [sname, setSName] = useState('');
  const [successMsg, setSuccessMsg] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);
  const [profileData, setProfileData] = useState(null);


  // Base64 to ArrayBuffer
  function bufferDecode(value) {
    return Uint8Array.from(atob(value), c => c.charCodeAt(0));
  }

  // ArrayBuffer to URLBase64
  function bufferEncode(value) {
    return btoa(String.fromCharCode.apply(null, new Uint8Array(value)))
      .replace(/\+/g, "-")
      .replace(/\//g, "_")
      .replace(/=/g, "");;
  }

  const registerName = () => {
    if (sname === "") {
      alert("Please enter a username");
      return;
    }
  }

  const accessName = () => {
    if (sname === "") {
      alert("Please enter a username");
      return;
    }
  }

  const handleUsernameChange = (e) => {
    setSName(e.target.value);
  };

  useEffect(() => {
    ;
    if (localStorage.getItem('loggedIn'))
      setLoggedIn(true);
    if (loggedIn)
      getProfile()
        .then(data => {
          setProfileData(data);
        })
        .catch(err => {
          setErrMsg(err.response.data);
          localStorage.removeItem('loggedIn');
        });
  }, [loggedIn]);

  return (
    <div className='App-header'>
      <Grid container textAlign='center' verticalAlign='middle'>
        <Grid.Column style={{ maxWidth: 450, minWidth: 300 }}>
          <Header as='h2' textAlign='center' style={{ color: 'white' }}>
            WebAuthn Demo
          </Header>
          {!loggedIn ?
            <Form size='large'>
              {errMsg && <Message negative icon='warning sign' size='mini' header={errMsg} />}
              {successMsg && <Message positive icon='thumbs up' size='mini' header={successMsg} />}
              <Segment>
                <Form.Input
                  fluid
                  icon='user'
                  iconPosition='left'
                  placeholder='Username'
                  onChange={handleUsernameChange}
                />
                <Button
                  fluid
                  size='large'
                  onClick={registerName}
                  style={{
                    marginTop: 8,
                    color: 'white',
                    backgroundColor: '#19857b'
                  }}
                  disabled={!sname}
                >
                  Register
                </Button>
                <LoginButton
                  domain="foo"
                  label="Login"
                  styling={"inline-flex items-center px-4 py-2 text-white bg-primaryLight-500 rounded hover:bg-primaryLight-700"}
                  onLogin={(result) => {
                    console.log(result)
                    setLoggedIn(result);
                  }}
                  onError={function (error) {
                    console.log(error);
                    setLoggedIn(false);
                  }}
                />
              </Segment>
            </Form>
            :
            <Segment style={{ overflowWrap: 'break-word' }}>
              {profileData &&
                <>
                  <Header as='h3' textAlign='center'>
                    Hi {profileData.name}
                  </Header>
                  <Header as='h4' textAlign='center'>
                    Your profile information
                  </Header>
                  <strong>ID: </strong>{profileData.id}
                  <br />
                  <strong>Credential information:</strong>
                  <br />
                  <strong>Format: </strong>{profileData.authenticators[0].fmt}
                  <br />
                  <strong>Public key: </strong>
                  <br />
                  <div style={{
                    maxWidth: 300,
                    overflowWrap: 'break-word',
                    marginLeft: '25%',
                    marginRight: '25%'
                  }}>
                    {profileData.authenticators[0].publicKey}
                  </div>
                  <Button
                    fluid
                    size='large'
                    // onClick={handleLogout}
                    style={{
                      marginTop: 8,
                      color: 'white',
                      backgroundColor: '#19857b'
                    }}
                  >
                    Logout
                  </Button>
                </>
              }
            </Segment>
          }
        </Grid.Column>
      </Grid>
    </div>
  );
}

export default App;
