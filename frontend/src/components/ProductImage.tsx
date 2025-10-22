import React from "react";

interface ImageSource {
  src: string;
  media: string;
}

interface ProductImageProps {
  src: string; // Default image source
  alt: string;
  sources?: ImageSource[]; // Optional responsive sources
}

const ProductImage: React.FC<ProductImageProps> = ({ src, alt, sources }) => {
  return (
    <picture>
      {sources?.map((source, index) => (
        <source key={index} srcSet={source.src} media={source.media} />
      ))}
      <img src={src} alt={alt} className="product-image" />
    </picture>
  );
};

export default ProductImage;
