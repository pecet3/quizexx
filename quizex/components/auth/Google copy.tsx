import React, { useState } from 'react';
import { View, TouchableOpacity, Text, StyleSheet } from 'react-native';
import * as WebBrowser from 'expo-web-browser';
import * as AuthSession from 'expo-auth-session';
import { makeRedirectUri } from 'expo-auth-session';
import Constants from 'expo-constants';
import * as SecureStore from 'expo-secure-store';
import { AntDesign } from "@expo/vector-icons"
// Konfiguracja do AuthSession
WebBrowser.maybeCompleteAuthSession();

const GoogleLogin = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<null | string>(null);

  // Pobierz apiUrl z konfiguracji
  const apiUrl = Constants.expoConfig?.extra?.apiUrl || 'https://c650-83-28-218-33.ngrok-free.app';

  // Konfiguracja URI przekierowania
  const redirectUri = makeRedirectUri({
    scheme: "myapp",
    path: "v1/google-callback",
  });


  // Funkcja do inicjowania procesu logowania
  const handleGoogleLogin = async () => {
    try {
      setLoading(true);
      setError(null);

      const authUrl = `${apiUrl}/v1/auth`;
      console.log('Auth URL:', authUrl);

      const result = await WebBrowser.openAuthSessionAsync(
        authUrl,
        redirectUri,
        {
          showInRecents: true,
        }
      );

      console.log('Auth result:', result); // dla debugowania

      if (result.type === 'success') {
        console.log("success")
        // Pobierz token z URL
        const token = extractTokenFromUrl(result.url);

        if (token) {
          // Zapisz token w secure storage
          await saveToken(token);
          // Nawiguj do głównego ekranu lub wykonaj inne akcje po zalogowaniu
          console.log('Token saved successfully');
        } else {
          setError('Nie udało się uzyskać tokenu');
        }
      } else if (result.type === 'cancel') {
        setError('Logowanie zostało anulowane');
      }
    } catch (err) {
      setError('Wystąpił błąd podczas logowania');
      console.error('Login error:', err);
    } finally {
      setLoading(false);
    }
  };

  // Helper do wyciągania tokenu z URL
  const extractTokenFromUrl = (url: any) => {
    try {
      const regex = /token=([^&]+)/;
      const match = url.match(regex);
      return match ? match[1] : null;
    } catch (err) {
      console.error('URL parsing error:', err);
      return null;
    }
  };

  // Funkcja do zapisywania tokenu
  const saveToken = async (token: any) => {
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