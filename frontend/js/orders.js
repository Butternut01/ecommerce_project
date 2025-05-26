import { getOrders } from './api.js';
import { getAuthToken, getCurrentUser } from './auth.js';

const orderList = document.getElementById('order-list');

// Load and display orders
export async function loadOrders() {
    const user = getCurrentUser();
    if (!user) {
        orderList.innerHTML = '<p>Please login to view your orders.</p>';
        return;
    }
    
    try {
        const orders = await getOrders(getAuthToken());
        displayOrders(orders.orders);
    } catch (error) {
        console.error('Failed to load orders:', error);
        orderList.innerHTML = '<p>Failed to load orders. Please try again later.</p>';
    }
}

function displayOrders(orders) {
    orderList.innerHTML = '';
    
    if (!orders || orders.length === 0) {
        orderList.innerHTML = '<p>You have no orders yet.</p>';
        return;
    }
    
    orders.forEach(order => {
        const orderElement = document.createElement('div');
        orderElement.className = 'order';
        orderElement.innerHTML = `
            <h3>Order #${order.id}</h3>
            <p>Status: ${order.status}</p>
            <p>Total: $${order.total.toFixed(2)}</p>
            <p>Date: ${new Date(order.created_at * 1000).toLocaleDateString()}</p>
            <div class="order-items">
                <h4>Items:</h4>
                <ul>
                    ${order.items.map(item => `
                        <li>
                            ${item.quantity} x Product ${item.product_id} 
                            ($${item.price.toFixed(2)} each)
                        </li>
                    `).join('')}
                </ul>
            </div>
        `;
        orderList.appendChild(orderElement);
    });
}