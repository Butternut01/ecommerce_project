import { registerUser, loginUser } from './api.js';

let currentUser = null;
let authToken = null;

// DOM elements
const loginBtn = document.getElementById('login-btn');
const registerBtn = document.getElementById('register-btn');
const logoutBtn = document.getElementById('logout-btn');
const userInfo = document.getElementById('user-info');
const loginForm = document.getElementById('login-form');
const registerForm = document.getElementById('register-form');
const loginFormData = document.getElementById('login-form-data');
const registerFormData = document.getElementById('register-form-data');

// Event listeners
loginBtn.addEventListener('click', () => {
    loginForm.style.display = 'block';
    registerForm.style.display = 'none';
});

registerBtn.addEventListener('click', () => {
    registerForm.style.display = 'block';
    loginForm.style.display = 'none';
});

logoutBtn.addEventListener('click', logout);

loginFormData.addEventListener('submit', async (e) => {
    e.preventDefault();
    const username = e.target[0].value;
    const password = e.target[1].value;
    
    try {
        const response = await loginUser({ username, password });
        authToken = response.token;
        currentUser = { id: response.user_id, username };
        updateAuthUI();
        loginForm.style.display = 'none';
        alert('Login successful!');
    } catch (error) {
        alert(`Login failed: ${error.message}`);
    }
});

registerFormData.addEventListener('submit', async (e) => {
    e.preventDefault();
    const username = e.target[0].value;
    const email = e.target[1].value;
    const password = e.target[2].value;
    
    try {
        await registerUser({ username, email, password });
        registerForm.style.display = 'none';
        alert('Registration successful! Please login.');
    } catch (error) {
        alert(`Registration failed: ${error.message}`);
    }
});

function updateAuthUI() {
    if (currentUser) {
        loginBtn.style.display = 'none';
        registerBtn.style.display = 'none';
        userInfo.style.display = 'inline';
        userInfo.textContent = `Welcome, ${currentUser.username}`;
        logoutBtn.style.display = 'inline';
    } else {
        loginBtn.style.display = 'inline';
        registerBtn.style.display = 'inline';
        userInfo.style.display = 'none';
        logoutBtn.style.display = 'none';
    }
}

function logout() {
    currentUser = null;
    authToken = null;
    updateAuthUI();
    alert('Logged out successfully');
}

// Check for existing session on page load
document.addEventListener('DOMContentLoaded', () => {
    // In a real app, you would check localStorage for a saved token
    updateAuthUI();
});

export function getAuthToken() {
    return authToken;
}

export function getCurrentUser() {
    return currentUser;
}