import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, StyleSheet, Alert } from 'react-native';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';

const API_URL = 'https://775b-89-64-30-38.ngrok-free.app'; // Replace with your actual API URL
const PREFIX = '/v1/auth'; // Replace with your API prefix

const AuthMagicCode = () => {
    const { user, setJWT } = useAuth();
    const [currentStep, setCurrentStep] = useState('register'); // register, login, exchange
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        code: '',
    });
    const [loading, setLoading] = useState(false);

    const handleRegister = async () => {
        try {
            setLoading(true);
            const response = await axios.post(`${API_URL}${PREFIX}/register`, {
                name: formData.name,
                email: formData.email,
            });
            if (response.status === 200) {
                Alert.alert('Success', 'Registration successful! Please check your email for verification code.');
                setCurrentStep('exchange');
            }
        } catch (error) {
        } finally {
            setLoading(false);
        }
    };

    const handleLogin = async () => {
        try {
            setLoading(true);
            const response = await axios.post(`${API_URL}${PREFIX}/login`, {
                email: formData.email,
            });

            if (response.status === 200) {
                Alert.alert('Success', 'Registration successful! Please check your email for verification code.');
                setCurrentStep('exchange');
            }
        } catch (error) {
            if (error !== null) {
            }
        } finally {
            setLoading(false);
        }
    };

    const handleExchange = async () => {
        try {
            setLoading(true);
            const response = await axios.post(`${API_URL}${PREFIX}/exchange`, {
                email: formData.email,
                code: formData.code,
            });

            if (response.data) {
                // Store the token or handle successful authentication
                Alert.alert('Success', 'Authentication successful!');
                console.log(response.data)
            }
        } catch (error) {
        } finally {
            setLoading(false);
        }
    };

    const renderRegisterForm = () => (
        <>
            <TextInput
                style={styles.input}
                placeholder="Name"
                value={formData.name}
                onChangeText={(text) => setFormData({ ...formData, name: text })}
            />
            <TextInput
                style={styles.input}
                placeholder="Email"
                value={formData.email}
                onChangeText={(text) => setFormData({ ...formData, email: text })}
                keyboardType="email-address"
                autoCapitalize="none"
            />
            <TouchableOpacity
                style={styles.button}
                onPress={handleRegister}
                disabled={loading}
            >
                <Text style={styles.buttonText}>
                    {loading ? 'Registering...' : 'Register'}
                </Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={() => setCurrentStep('login')}>
                <Text style={styles.link}>Already have an account? Login</Text>
            </TouchableOpacity>
        </>
    );

    const renderLoginForm = () => (
        <>
            <TextInput
                style={styles.input}
                placeholder="Email"
                value={formData.email}
                onChangeText={(text) => setFormData({ ...formData, email: text })}
                keyboardType="email-address"
                autoCapitalize="none"
            />
            <TouchableOpacity
                style={styles.button}
                onPress={handleLogin}
                disabled={loading}
            >
                <Text style={styles.buttonText}>
                    {loading ? 'Logging in...' : 'Login'}
                </Text>
            </TouchableOpacity>
            <TouchableOpacity onPress={() => setCurrentStep('register')}>
                <Text style={styles.link}>Don't have an account? Register</Text>
            </TouchableOpacity>
        </>
    );

    const renderExchangeForm = () => (
        <>
            <TextInput
                style={styles.input}
                placeholder="Verification Code"
                value={formData.code}
                onChangeText={(text) => setFormData({ ...formData, code: text })}
                keyboardType="numeric"
                maxLength={6}
            />
            <TouchableOpacity
                style={styles.button}
                onPress={handleExchange}
                disabled={loading}
            >
                <Text style={styles.buttonText}>
                    {loading ? 'Verifying...' : 'Verify Code'}
                </Text>
            </TouchableOpacity>
        </>
    );

    return (
        <View style={styles.container}>
            <Text style={styles.title}>
                {currentStep === 'register' ? 'Register' :
                    currentStep === 'login' ? 'Login' : 'Verify Code'}
            </Text>

            {currentStep === 'register' && renderRegisterForm()}
            {currentStep === 'login' && renderLoginForm()}
            {currentStep === 'exchange' && renderExchangeForm()}
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 20,
        justifyContent: 'center',
        backgroundColor: '#fff',
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 20,
        textAlign: 'center',
    },
    input: {
        height: 50,
        borderWidth: 1,
        borderColor: '#ddd',
        borderRadius: 8,
        paddingHorizontal: 15,
        marginBottom: 15,
        fontSize: 16,
    },
    button: {
        backgroundColor: '#007AFF',
        padding: 15,
        borderRadius: 8,
        marginBottom: 15,
    },
    buttonText: {
        color: '#fff',
        textAlign: 'center',
        fontSize: 16,
        fontWeight: 'bold',
    },
    link: {
        color: '#007AFF',
        textAlign: 'center',
        fontSize: 14,
        marginTop: 10,
    },
});

export default AuthMagicCode;