import React from "react";
import { useCartDispatch } from "../context/CartContext";

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  category: string;
}

interface AddToCartButtonProps {
  product: Product;
}

const AddToCartButton: React.FC<AddToCartButtonProps> = ({ product }) => {
  const dispatch = useCartDispatch();

  const handleAddToCart = () => {
    dispatch({ type: "ADD_ITEM", payload: product });
  };

  return (
    <button onClick={handleAddToCart} className="add-to-cart-button">
      Add to Cart
    </button>
  );
};

export default AddToCartButton;
