import React from "react";
import { useCartDispatch } from "../context/CartContext";

interface Product {
  id: string;
  name: string;
  price: number;
  imageUrl: string;
}

interface CartItemProps {
  item: {
    product: Product;
    quantity: number;
  };
}

const CartItem: React.FC<CartItemProps> = ({ item }) => {
  const dispatch = useCartDispatch();

  const handleIncreaseQuantity = () => {
    dispatch({ type: "INCREASE_QUANTITY", payload: item.product.id });
  };

  const handleDecreaseQuantity = () => {
    dispatch({ type: "DECREASE_QUANTITY", payload: item.product.id });
  };

  const handleRemoveItem = () => {
    dispatch({ type: "REMOVE_ITEM", payload: item.product.id });
  };

  return (
    <div className="cart-item">
      <img
        src={item.product.imageUrl}
        alt={item.product.name}
        className="cart-item-image"
      />
      <div className="cart-item-details">
        <h4 className="cart-item-name">{item.product.name}</h4>
        <p className="cart-item-price">${item.product.price.toFixed(2)}</p>
        <div className="cart-item-quantity-controls">
          <button onClick={handleDecreaseQuantity}>-</button>
          <span>{item.quantity}</span>
          <button onClick={handleIncreaseQuantity}>+</button>
        </div>
        <button onClick={handleRemoveItem} className="cart-item-remove-button">
          Remove
        </button>
      </div>
    </div>
  );
};

export default CartItem;
