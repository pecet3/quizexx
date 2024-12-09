import React, { createContext, useContext, useState, ReactNode, useEffect } from 'react';
import { BasketItem, Price, Product } from './types';


type StoreContextType = {
    isLoading: boolean
    setIsLoading: React.Dispatch<React.SetStateAction<boolean>>
    products: Product[]
    addProducts: (products: Product[]) => void
    getProductById: (id: number) => Product | undefined
    prices: Price[]
    addPrices: (prices: Price[]) => void
    basket: BasketItem[]
    addItemToBasket: (newItem: BasketItem) => void
    clearBasket: () => void
    changeItemQuantity: (item: BasketItem, quantity: number) => void
    itemsCount: number
    finalPrice: number
}

const StoreContext = createContext<StoreContextType | undefined>(undefined);

export const StoreProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [isLoading, setIsLoading] = useState(true)
    const [products, setProducts] = useState<Product[]>([{
        id: 0,
        name: "",
        image_url: "",
        description: "",
    }]);
    const addProducts = (products: Product[]) => {
        setProducts(products);
    };
    const getProductById = (id: number) => {
        return products.find(p => p.id === id)
    }

    const [prices, setPrices] = useState<Price[]>([{
        id: 0,
        price: 0,
        capacity: 0
    }])
    const addPrices = (prices: Price[]) => {
        setPrices(prices)
    }

    const [basket, setBasket] = useState<BasketItem[]>(() => {
        const savedBasket = localStorage.getItem('basket')
        return savedBasket ? JSON.parse(savedBasket) : []
    })
    useEffect(() => {
        localStorage.setItem('basket', JSON.stringify(basket))
    }, [basket])
    const clearBasket = () => {
        setBasket([])
    }
    const addItemToBasket = (newItem: BasketItem) => {
        const existingItem = basket.find(item =>
            item.product.id === newItem.product.id
            && item.capacity === newItem.capacity
            && item.strength === newItem.strength)

        if (existingItem) {
            setBasket(prev => prev.map(i =>
                i.product.id === existingItem.product.id
                    && i.capacity === existingItem.capacity
                    && i.strength === existingItem.strength
                    ? { ...i, quantity: i.quantity + newItem.quantity } : i))
            return
        }
        setBasket(prev => [...prev, newItem])
        return
    }
    const changeItemQuantity = (item: BasketItem, quantity: number) => {
        const existingItem = basket.find(i =>
            i.product.id === item.product.id
            && i.capacity === item.capacity
            && i.strength === item.strength)
        if (!existingItem) return
        if (quantity <= 0) {
            setBasket(basket.filter(i => i !== existingItem))
        }
        setBasket(prev => prev.map(i =>
            i.product.id === existingItem.product.id
                && i.capacity === existingItem.capacity
                && i.strength === existingItem.strength
                ? { ...i, quantity: quantity } : i))
        return
    }
    const itemsCount = ((): number => {
        let c = 0
        for (let item of basket) {
            c = c + item.quantity
        }
        return c
    })()
    const finalPrice = Number(((): number => {
        let p = 0
        for (let item of basket) {
            p = p + item.price * item.quantity
        }
        return p
    })().toFixed(2))
    return (
        <StoreContext.Provider value={
            {
                isLoading, setIsLoading,
                products, addProducts, getProductById,
                prices, addPrices,
                basket, addItemToBasket, clearBasket, itemsCount, finalPrice, changeItemQuantity
            }}>
            {children}
        </StoreContext.Provider>
    );
};

export const useStoreContext = () => {
    const context = useContext(StoreContext);
    if (context === undefined) {
        throw new Error('useStoreContext must be used within a StoreProvider');
    }
    return context;
};