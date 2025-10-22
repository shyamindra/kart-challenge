import React from "react";
import ProductImage from "./ProductImage";
import AddToCartButton from "./AddToCartButton";

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  category: string;
}

interface ProductCardProps {
  product: Product;
}

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  return (
    <div className="product-card">
      <ProductImage src={product.imageUrl} alt={product.name} />
      <div className="product-info">
        <h3 className="product-name">{product.name}</h3>
        <p className="product-category">{product.category}</p>
        <p className="product-price">${product.price.toFixed(2)}</p>
        <AddToCartButton product={product} />
      </div>
    </div>
  );
};

export default ProductCard;
