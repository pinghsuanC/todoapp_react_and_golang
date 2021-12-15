import React from "react";
import logo from "./logo.svg";
import "./App.css";
import { Header } from "./Components/Header";
import { TodoList } from "./Components/TodoList";

function App() {
	return (
		<div className="App">
			<Header />
			<TodoList
				todos={[
					{
						title: "Do Dishes",
						description: "Wash the dishes by hand",
						isCompleted: true,
					},
					{ title: "Do Laundry", isCompleted: false },
					{
						title: "Do Homework",
						description: "Finish English reading",
						isCompleted: false,
					},
				]}
			/>
		</div>
	);
}

export default App;
