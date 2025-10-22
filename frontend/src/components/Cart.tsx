import React, { useState } from "react";
import { useCart, useCartDispatch } from "../context/CartContext";
import CartItem from "./CartItem";
import { placeOrder } from "../services/apiService";
import OrderConfirmation from "./OrderConfirmation";

const Cart: React.FC = () => {
  const { items } = useCart();
  const dispatch = useCartDispatch();
  const [showConfirmation, setShowConfirmation] = useState(false);
  const [confirmedOrder, setConfirmedOrder] = useState<{
    orderId: string;
    totalAmount: number;
  } | null>(null);
  const [couponCode, setCouponCode] = useState<string>("");

  const calculateTotal = () => {
    return items.reduce(
      (total, item) => total + item.product.price * item.quantity,
      0,
    );
  };

  const handleConfirmOrder = async () => {
    if (items.length === 0) {
      alert("Your cart is empty!");
      return;
    }

    const orderRequest = {
      items: items.map((item) => ({
        productId: item.product.id,
        quantity: item.quantity,
      })),
      couponCode: couponCode || undefined, // Pass coupon code if not empty
    };

    try {
      const orderResponse = await placeOrder(orderRequest);
      setConfirmedOrder({
        orderId: orderResponse.orderId,
        totalAmount: orderResponse.totalAmount,
      });
      setShowConfirmation(true);
      dispatch({ type: "CLEAR_CART" });
      setCouponCode(""); // Clear coupon code after successful order
    } catch (error) {
      let errorMessage = "An unknown error occurred.";
      if (error instanceof Error) {
        errorMessage = error.message;
      }
      alert(`Failed to place order: ${errorMessage}`);
      console.error("Error placing order:", error);
    }
  };

  const handleCloseConfirmation = () => {
    setShowConfirmation(false);
    setConfirmedOrder(null);
  };

  return (
    <div className="cart">
      <h2>Your Cart</h2>
      {items.length === 0 ? (
        <p>Your cart is empty.</p>
      ) : (
        <div className="cart-items-list">
          {items.map((item) => (
            <CartItem key={item.product.id} item={item} />
          ))}
        </div>
      )}
      <div className="cart-summary">
        <div className="coupon-input">
          <input
            type="text"
            placeholder="Coupon Code"
            value={couponCode}
            onChange={(e) => setCouponCode(e.target.value)}
          />
        </div>
        <p>Total: ${calculateTotal().toFixed(2)}</p>
        <button
          onClick={handleConfirmOrder}
          className="confirm-order-button"
          disabled={items.length === 0}
        >
          Confirm Order
        </button>
      </div>

      {showConfirmation && confirmedOrder && (
        <OrderConfirmation
          orderId={confirmedOrder.orderId}
          totalAmount={confirmedOrder.totalAmount}
          onClose={handleCloseConfirmation}
        />
      )}
    </div>
  );
};

export default Cart;
