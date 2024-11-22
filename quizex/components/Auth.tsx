import React, { useState } from "react";
import { View, TouchableOpacity, Text, StyleSheet } from "react-native";
import * as WebBrowser from "expo-web-browser";
import * as AuthSession from "expo-auth-session";
import { makeRedirectUri } from "expo-auth-session";
import Constants from "expo-constants";
import * as SecureStore from "expo-secure-store";

// Konfiguracja do AuthSession
WebBrowser.maybeCompleteAuthSession();

const GoogleLogin = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  // Konfiguracja URI przekierowania
  const redirectUri = makeRedirectUri({
    scheme: "your-app-scheme", // Zmień na scheme twojej aplikacji
    path: "google-callback",
  });

  // Funkcja do inicjowania procesu logowania
  const handleGoogleLogin = async () => {
    try {
      setLoading(true);
      setError("");

      // Rozpocznij proces autoryzacji
      const authUrl = `${Constants.manifest.extra.apiUrl}/v1/auth`;
      const result = await WebBrowser.openAuthSessionAsync(
        authUrl,
        redirectUri,
        {
          showInRecents: true,
        }
      );

      if (result.type === "success") {
        // Pobierz token z URL
        const token = extractTokenFromUrl(result.url);

        if (token) {
          // Zapisz token w secure storage
          await saveToken(token);
          // Nawiguj do głównego ekranu lub wykonaj inne akcje po zalogowaniu
        } else {
          setError("Nie udało się uzyskać tokenu");
        }
      }
    } catch (err) {
      setError("Wystąpił błąd podczas logowania");
      console.error(err);
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
      console.error("Błąd podczas parsowania URL:", err);
      return null;
    }
  };

  // Funkcja do zapisywania tokenu
  const saveToken = async (token: any) => {
    try {
      await SecureStore.setItemAsync("userToken", token);
    } catch (err) {
      console.error("Błąd podczas zapisywania tokenu:", err);
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
        <Text style={styles.buttonText}>
          {loading ? "Logowanie..." : "Zaloguj się przez Google"}
        </Text>
      </TouchableOpacity>

      {error && <Text style={styles.errorText}>{error}</Text>}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    padding: 20,
  },
  googleButton: {
    backgroundColor: "#4285F4",
    paddingHorizontal: 20,
    paddingVertical: 12,
    borderRadius: 5,
    width: "100%",
    maxWidth: 300,
  },
  buttonText: {
    color: "white",
    fontSize: 16,
    textAlign: "center",
    fontWeight: "600",
  },
  errorText: {
    color: "red",
    marginTop: 10,
    textAlign: "center",
  },
});

export default GoogleLogin;
