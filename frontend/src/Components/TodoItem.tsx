import React from "react";
import { Todo } from "../entities/Todo";
import "./TodoItem.css";

const TodoItem: React.FC<Todo> = ({ title, description, isCompleted }) => {
	return (
		<article className="todo-item">
			<section className="todo-item-text">
				<h3>{title}</h3>
				<p>{description}</p>
			</section>
			<input type="checkbox" defaultChecked={isCompleted} />
		</article>
	);
};

export default TodoItem;
