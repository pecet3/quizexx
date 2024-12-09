import React, { useState, useEffect } from 'react';
import { View, Button, ActivityIndicator, Text, TouchableOpacity, StyleSheet } from 'react-native';
import * as WebBrowser from 'expo-web-browser';
import { makeRedirectUri } from 'expo-auth-session';
import * as Linking from 'expo-linking';
import * as SecureStore from 'expo-secure-store';
import Constants from 'expo-constants';
import { AntDesign } from '@expo/vector-icons';
import * as Crypto from 'expo-crypto';

const GoogleLogin = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [pubCode, setPubCode] = useState("")
  const [secretCode, setSecretCode] = useState("")
  const apiUrl = Constants.expoConfig?.extra?.apiUrl || 'https://267a-89-64-30-38.ngrok-free.app';

  // Konfiguracja URI przekierowania
  const redirectUri = makeRedirectUri({
    scheme: "myapp",
    path: "callback/google",
    preferLocalhost: true
  });


  // Inicjacja procesu logowania
  const handleGoogleLogin = async () => {
    try {
      setLoading(true);
      setError(null);
      const pubCode = Crypto.randomUUID()
      console.log(pubCode)
      setPubCode(pubCode)

      const secretCodeResult = await fetch(`${apiUrl}/v1/exchange?pubCode=${pubCode}`)
      const secretCode = await secretCodeResult.json()
      console.log(secretCode)
      setSecretCode(secretCode)
      const authUrl = `${apiUrl}/v1/auth?pubCode=${pubCode}`;
      console.log('Auth URL:', authUrl);

      await WebBrowser.maybeCompleteAuthSession();
      const result = await WebBrowser.openAuthSessionAsync(
        authUrl,
        redirectUri,
        {
          showInRecents: false,
          preferEphemeralSession: true // Używaj sesji tymczasowej
        }
      );

      console.log('Auth result:', result);

      if (result.type === 'dismiss') {
        setError('Logowanie zostało anulowane');
      }
    } catch (err) {
      console.error('Login error:', err);
      setError('Wystąpił błąd podczas logowania');
    } finally {
      setLoading(false);
    }
  };

  const extractTokenFromUrl = (url: string) => {
    try {
      const regex = /token=([^&]+)/;
      const match = url.match(regex);
      return match ? match[1] : null;
    } catch (err) {
      console.error('URL parsing error:', err);
      return null;
    }
  };

  const saveToken = async (token: string) => {
    try {
      await SecureStore.setItemAsync('userToken', token);
    } catch (err) {
      console.error('Token save error:', err);
      throw err;
    }
  };


  return (
    <View style={styles.container}>
      <TouchableOpacity
        style={styles.googleButton}
        onPress={handleGoogleLogin}
        disabled={loading}
      >
        <View style={styles.buttonContent}>
          <AntDesign
            name="google"
            size={20}
            color="white"
            style={styles.icon}
          />
          <Text style={styles.buttonText}>
            {loading ? 'Loading...' : 'oogle'}
          </Text>
        </View>
      </TouchableOpacity>

      {error && (
        <Text style={styles.errorText}>{error}</Text>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  googleButton: {
    backgroundColor: '#4285F4',
    paddingHorizontal: 20,
    paddingVertical: 12,
    borderRadius: 5,
    width: '100%',
    maxWidth: 300,
  },
  buttonContent: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
  },
  buttonText: {
    color: 'white',
    fontSize: 20,
    fontWeight: '600',
    marginRight: 10,
  },
  icon: {
    marginLeft: 0,
  },
  errorText: {
    color: 'red',
    marginTop: 10,
    textAlign: 'center',
  },
});
export default GoogleLogin;