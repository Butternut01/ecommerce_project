const API_BASE_URL = 'http://localhost:8000'; // API Gateway URL

// Generic API call function
async function callApi(endpoint, method = 'GET', data = null, token = null) {
    const url = `${API_BASE_URL}${endpoint}`;
    const headers = {
        'Content-Type': 'application/json',
    };
    
    if (token) {
        headers['Authorization'] = `Bearer ${token}`;
    }
    
    const options = {
        method,
        headers,
    };
    
    if (data) {
        options.body = JSON.stringify(data);
    }
    
    try {
        const response = await fetch(url, options);
        const responseData = await response.json();
        
        if (!response.ok) {
            throw new Error(responseData.error || 'Something went wrong');
        }
        
        return responseData;
    } catch (error) {
        console.error('API call failed:', error);
        throw error;
    }
}

// Auth API functions
export async function registerUser(userData) {
    return callApi('/users/register', 'POST', userData);
}

export async function loginUser(credentials) {
    return callApi('/users/login', 'POST', credentials);
}

// Product API functions
export async function getProducts() {
    const response = await fetch('http://localhost:8080/products', {
        credentials: 'include'
    });

    if (!response.ok) {
        throw new Error('Failed to fetch products');
    }
    return await response.json(); // should be an array []
}

export async function getProduct(id) {
    return callApi(`/products/${id}`);
}

// Order API functions
export async function getOrders(token) {
    return callApi('/orders', 'GET', null, token);
}

export async function createOrder(orderData, token) {
    return callApi('/orders', 'POST', orderData, token);
}