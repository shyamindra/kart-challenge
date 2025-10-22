import React from "react";

interface OrderConfirmationProps {
  orderId: string;
  totalAmount: number;
  onClose: () => void;
}

const OrderConfirmation: React.FC<OrderConfirmationProps> = ({
  orderId,
  totalAmount,
  onClose,
}) => {
  return (
    <div className="order-confirmation-modal">
      <div className="modal-content">
        <h2>Order Confirmed!</h2>
        <p>Your order has been placed successfully.</p>
        <p>
          Order ID: <strong>{orderId}</strong>
        </p>
        <p>
          Total Amount: <strong>${totalAmount.toFixed(2)}</strong>
        </p>
        <button onClick={onClose}>Close</button>
      </div>
    </div>
  );
};

export default OrderConfirmation;
