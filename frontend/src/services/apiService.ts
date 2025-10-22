const API_BASE_URL = "http://localhost:8080"; // Assuming backend runs on 8080
const API_KEY = "your_api_key_here"; // Placeholder for API Key

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  category: string;
}

interface OrderItem {
  productId: string;
  quantity: number;
}

interface OrderRequest {
  items: OrderItem[];
  couponCode?: string;
}

interface OrderResponse {
  orderId: string;
  totalAmount: number;
  // Add other fields as per backend API response
}

export const getProducts = async (): Promise<Product[]> => {
  try {
    const response = await fetch(`${API_BASE_URL}/product`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data: Product[] = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching products:", error);
    throw error;
  }
};

export const placeOrder = async (
  orderData: OrderRequest,
): Promise<OrderResponse> => {
  try {
    const response = await fetch(`${API_BASE_URL}/order`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-API-Key": API_KEY, // Sending API key in header as per plan
      },
      body: JSON.stringify(orderData),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(
        errorData.message || `HTTP error! status: ${response.status}`,
      );
    }

    const data: OrderResponse = await response.json();
    return data;
  } catch (error) {
    console.error("Error placing order:", error);
    throw error;
  }
};
