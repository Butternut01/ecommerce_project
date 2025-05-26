import { getProducts, getProduct } from './api.js';
import { getAuthToken, getCurrentUser } from './auth.js';

const productList = document.getElementById('product-list');
const productModal = document.getElementById('product-modal');
const productModalTitle = document.getElementById('product-modal-title');
const productModalContent = document.getElementById('product-modal-content');
const productOrderForm = document.getElementById('product-order-form');
const orderForm = document.getElementById('order-form');
const closeModal = document.querySelector('.close');

// Load and display products
export async function loadProducts() {
    try {
        const products = await getProducts();
        displayProducts(products); // ✅ Pass array directly
    } catch (error) {
        console.error('Failed to load products:', error);
        productList.innerHTML = '<p>Failed to load products. Please try again later.</p>';
    }
}


function displayProducts(products) {
    productList.innerHTML = '';
    
    if (!products || products.length === 0) {
        productList.innerHTML = '<p>No products available.</p>';
        return;
    }
    
    products.forEach(product => {
        const productCard = document.createElement('div');
        productCard.className = 'product-card';
        productCard.innerHTML = `
            <h3>${product.name}</h3>
            <p>${product.description}</p>
            <p>Price: $${product.price.toFixed(2)}</p>
            <p>Stock: ${product.stock}</p>
            <button class="view-product" data-id="${product.id}">View Details</button>
        `;
        productList.appendChild(productCard);
    });
    
    // Add event listeners to view buttons
    document.querySelectorAll('.view-product').forEach(button => {
        button.addEventListener('click', () => viewProductDetails(button.getAttribute('data-id')));
    });
}

async function viewProductDetails(productId) {
    try {
        const product = await getProduct(productId);
        showProductModal(product);
    } catch (error) {
        console.error('Failed to load product details:', error);
        alert('Failed to load product details. Please try again later.');
    }
}

function showProductModal(product) {
    productModalTitle.textContent = product.name;
    productModalContent.innerHTML = `
        <p><strong>Description:</strong> ${product.description}</p>
        <p><strong>Price:</strong> $${product.price.toFixed(2)}</p>
        <p><strong>Stock:</strong> ${product.stock}</p>
        <p><strong>Category:</strong> ${product.category}</p>
    `;
    
    // Show order form if user is logged in
    if (getCurrentUser()) {
        productOrderForm.style.display = 'block';
        orderForm.dataset.productId = product.id;
    } else {
        productOrderForm.style.display = 'none';
    }
    
    productModal.style.display = 'block';
}

// Close modal when clicking X
closeModal.addEventListener('click', () => {
    productModal.style.display = 'none';
});

// Close modal when clicking outside
window.addEventListener('click', (e) => {
    if (e.target === productModal) {
        productModal.style.display = 'none';
    }
});

// Handle order form submission
orderForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const productId = e.target.dataset.productId;
    const quantity = parseInt(document.getElementById('order-quantity').value);
    
    // In a real app, you would call createOrder here
    alert(`Order placed for ${quantity} of product ${productId}`);
    productModal.style.display = 'none';
});