import React, { createContext, useReducer, useContext, ReactNode } from "react";

// Reusing Product interface from apiService.ts for consistency
interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  category: string;
}

interface CartItem {
  product: Product;
  quantity: number;
}

interface CartState {
  items: CartItem[];
}

type CartAction =
  | { type: "ADD_ITEM"; payload: Product }
  | { type: "REMOVE_ITEM"; payload: string } // payload is productId
  | { type: "INCREASE_QUANTITY"; payload: string } // payload is productId
  | { type: "DECREASE_QUANTITY"; payload: string } // payload is productId
  | { type: "CLEAR_CART" };

const initialCartState: CartState = {
  items: [],
};

const cartReducer = (state: CartState, action: CartAction): CartState => {
  switch (action.type) {
    case "ADD_ITEM": {
      const existingItemIndex = state.items.findIndex(
        (item) => item.product.id === action.payload.id,
      );

      if (existingItemIndex > -1) {
        const updatedItems = [...state.items];
        updatedItems[existingItemIndex].quantity += 1;
        return { ...state, items: updatedItems };
      } else {
        return {
          ...state,
          items: [...state.items, { product: action.payload, quantity: 1 }],
        };
      }
    }
    case "REMOVE_ITEM":
      return {
        ...state,
        items: state.items.filter((item) => item.product.id !== action.payload),
      };
    case "INCREASE_QUANTITY": {
      const updatedItems = state.items.map((item) =>
        item.product.id === action.payload
          ? { ...item, quantity: item.quantity + 1 }
          : item,
      );
      return { ...state, items: updatedItems };
    }
    case "DECREASE_QUANTITY": {
      const updatedItems = state.items
        .map((item) =>
          item.product.id === action.payload
            ? { ...item, quantity: Math.max(1, item.quantity - 1) } // Ensure quantity doesn't go below 1
            : item,
        )
        .filter((item) => item.quantity > 0); // Remove if quantity becomes 0
      return { ...state, items: updatedItems };
    }
    case "CLEAR_CART":
      return initialCartState;
    default:
      return state;
  }
};

const CartContext = createContext<CartState | undefined>(undefined);
const CartDispatchContext = createContext<
  React.Dispatch<CartAction> | undefined
>(undefined);

interface CartProviderProps {
  children: ReactNode;
}

export const CartProvider: React.FC<CartProviderProps> = ({ children }) => {
  const [state, dispatch] = useReducer(cartReducer, initialCartState);

  return (
    <CartContext.Provider value={state}>
      <CartDispatchContext.Provider value={dispatch}>
        {children}
      </CartDispatchContext.Provider>
    </CartContext.Provider>
  );
};

export const useCart = () => {
  const context = useContext(CartContext);
  if (context === undefined) {
    throw new Error("useCart must be used within a CartProvider");
  }
  return context;
};

export const useCartDispatch = () => {
  const context = useContext(CartDispatchContext);
  if (context === undefined) {
    throw new Error("useCartDispatch must be used within a CartProvider");
  }
  return context;
};
