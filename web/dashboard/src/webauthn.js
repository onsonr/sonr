
import axios from 'axios';
axios.defaults.withCredentials = true;

function getMakeCredentialsChallenge(formBody) {
    return axios.post('webauthn/register', formBody)
        .then(response => {
            if (response.data.status !== 'ok')
                throw new Error(`Server responed with error. The message is: ${response.message}`);
            return response.data;
        });
}

function sendWebAuthnResponse(body) {
    return axios.post('webauthn/response', body)
        .then(response => {
            if (response.data.status !== 'ok')
                throw new Error(`Server responed with error. The message is: ${response.message}`);
            return response.data;
        });
}

function getGetAssertionChallenge(formBody) {
    return axios.post('webauthn/login', formBody)
        .then(response => {
            if (response.data.status !== 'ok')
                throw new Error(`Server responed with error. The message is: ${response.message}`);
            return response.data;
        });
};

function getProfile() {
    return axios.get('webauthn/profile')
        .then(response => response.data);
}

function logout() {
    return axios.get('webauthn/profile')
        .then(response => response.data);
}
function registerFail(body) {
    return axios.post('webauthn/registerfail', body)
        .then(response => response.data);
}


export {
    getGetAssertionChallenge,
    getMakeCredentialsChallenge,
    sendWebAuthnResponse,
    getProfile,
    logout,
    registerFail
};
