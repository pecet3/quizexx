import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  ScrollView,
  StyleSheet,
  Platform,
} from "react-native";
import { Picker } from "@react-native-picker/picker";
import { StatusBar } from "expo-status-bar";

const CreateRoom = () => {
  const [roomName, setRoomName] = useState("");
  const [category, setCategory] = useState("");
  const [difficulty, setDifficulty] = useState("easy");
  const [maxRounds, setMaxRounds] = useState("4");
  const [language, setLanguage] = useState("polish");

  const handleSubmit = () => {
    // Handle form submission logic here
    console.log({
      roomName,
      category,
      difficulty,
      maxRounds,
      language,
    });
  };

  const getRandomCategory = () => {
    // Implement random category logic here
    console.log("Getting random category...");
  };

  return (
    <ScrollView style={styles.container}>
      <StatusBar style="auto" />

      {/* Main Form */}
      <View style={styles.paper}>
        <TextInput
          style={styles.input}
          placeholder="Room Name"
          value={roomName}
          onChangeText={setRoomName}
          placeholderTextColor="#666"
        />

        <View style={styles.categoryContainer}>
          <TextInput
            style={styles.input}
            placeholder="Category of Questions"
            value={category}
            onChangeText={setCategory}
            placeholderTextColor="#666"
          />

          <TouchableOpacity
            onPress={getRandomCategory}
            style={styles.randomButton}
          >
            <Text style={styles.randomButtonText}>[Get random category]</Text>
          </TouchableOpacity>

          <Text style={styles.categoryInfo}>
            Category can be anything, {"\n"}
            <Text style={styles.bold}>
              Quizex is connected with Chat-GPT-3.5
            </Text>
            .{"\n"}
            Based on the provided category, questions are prepared.
          </Text>
        </View>

        <Text style={styles.label}>Difficulty Level:</Text>
        <View style={styles.pickerContainer}>
          <Picker
            selectedValue={difficulty}
            onValueChange={setDifficulty}
            style={styles.picker}
          >
            <Picker.Item label="Easy" value="easy" />
            <Picker.Item label="Medium" value="medium" />
            <Picker.Item label="Hard" value="hard" />
          </Picker>
        </View>

        <View style={styles.rowContainer}>
          <View style={styles.pickerWrapper}>
            <Text style={styles.label}>Rounds:</Text>
            <View style={styles.pickerContainer}>
              <Picker
                selectedValue={maxRounds}
                onValueChange={setMaxRounds}
                style={styles.picker}
              >
                <Picker.Item label="4" value="4" />
                <Picker.Item label="5" value="5" />
                <Picker.Item label="6" value="6" />
              </Picker>
            </View>
          </View>

          <View style={styles.pickerWrapper}>
            <Text style={styles.label}>Language:</Text>
            <View style={styles.pickerContainer}>
              <Picker
                selectedValue={language}
                onValueChange={setLanguage}
                style={styles.picker}
              >
                <Picker.Item label="Polish" value="polish" />
                <Picker.Item label="English" value="english" />
              </Picker>
            </View>
          </View>
        </View>

        <TouchableOpacity style={styles.submitButton} onPress={handleSubmit}>
          <Text style={styles.submitButtonText}>Create Room</Text>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#9ca3af",
    padding: Platform.OS === "ios" ? 20 : 10,
  },
  header: {
    padding: 4,
    marginVertical: 16,
    alignItems: "center",
  },
  emoji: {
    fontSize: 24,
    position: "absolute",
    top: 0,
  },
  title: {
    fontSize: 48,
    fontWeight: "900",
    textDecorationLine: "underline",
    fontFamily: Platform.OS === "ios" ? "Courier" : "monospace",
  },
  paper: {
    backgroundColor: "#feffc2",
    borderRadius: 8,
    padding: 16,
    marginVertical: 16,
    maxWidth: 500,
    alignSelf: "center",
    width: "100%",
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    elevation: 5,
  },
  input: {
    backgroundColor: "white",
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 4,
    padding: 8,
    fontSize: 20,
    textAlign: "center",
    marginVertical: 8,
  },
  categoryContainer: {
    alignItems: "center",
    marginVertical: 8,
  },
  randomButton: {
    padding: 8,
  },
  randomButtonText: {
    fontSize: 14,
    color: "#666",
  },
  categoryInfo: {
    fontFamily: Platform.OS === "ios" ? "Courier" : "monospace",
    fontSize: 16,
    textAlign: "center",
    marginVertical: 8,
  },
  bold: {
    fontWeight: "bold",
    textDecorationLine: "underline",
  },
  label: {
    fontSize: 18,
    fontWeight: "bold",
    textDecorationLine: "underline",
    fontFamily: Platform.OS === "ios" ? "Courier" : "monospace",
    textAlign: "center",
    marginVertical: 8,
  },
  pickerContainer: {
    backgroundColor: "white",
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 4,
    marginVertical: 8,
  },
  picker: {
    height: 50,
    width: Platform.OS === "ios" ? 200 : "100%",
  },
  rowContainer: {
    flexDirection: "row",
    justifyContent: "space-around",
    flexWrap: "wrap",
    marginVertical: 8,
  },
  pickerWrapper: {
    alignItems: "center",
    marginHorizontal: 8,
  },
  submitButton: {
    backgroundColor: "#5eead4",
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 8,
    padding: 8,
    marginVertical: 16,
    alignSelf: "center",
  },
  submitButtonText: {
    fontSize: 20,
    fontFamily: Platform.OS === "ios" ? "Courier" : "monospace",
    fontWeight: "300",
  },
});

export default CreateRoom;
