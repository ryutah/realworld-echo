import {
  Dispatch,
  ReactNode,
  createContext,
  useContext,
  useReducer,
  useState,
} from "react";

const TaskContext = createContext<TaskType[]>([]);

type TaskDispatchType = Dispatch<TaskActions>;

const TaskDispatchContext = createContext<TaskDispatchType | undefined>(
  undefined
);

const useTasks = () => {
  return useContext(TaskContext);
};

const useTaskDispatch = () => {
  return useContext(TaskDispatchContext);
};

const TaskProvider = ({ children }: { children: ReactNode }) => {
  const [tasks, dispatch] = useReducer(taskReducer, initialTasks);

  return (
    <TaskContext.Provider value={tasks}>
      <TaskDispatchContext.Provider value={dispatch}>
        {children}
      </TaskDispatchContext.Provider>
    </TaskContext.Provider>
  );
};

const TaskActionType = {
  ADD_TASK: "task/add",
  CHANGE_TASK: "task/change",
  DELETE_TASK: "task/delete",
} as const;
type TaskActionType = (typeof TaskActionType)[keyof typeof TaskActionType];

type TaskActions =
  | {
      type: typeof TaskActionType.ADD_TASK;
      payload: TaskType;
    }
  | {
      type: typeof TaskActionType.CHANGE_TASK;
      payload: TaskType;
    }
  | {
      type: typeof TaskActionType.DELETE_TASK;
      payload: {
        id: number;
      };
    };

const taskReducer = (tasks: TaskType[], action: TaskActions) => {
  switch (action.type) {
    case TaskActionType.ADD_TASK:
      return [
        ...tasks,
        {
          id: action.payload.id,
          text: action.payload.text,
          done: action.payload.done,
        },
      ];
    case TaskActionType.CHANGE_TASK:
      return tasks.map((t) =>
        t.id === action.payload.id ? action.payload : t
      );
    case TaskActionType.DELETE_TASK:
      return tasks.filter((task) => task.id !== action.payload.id);
  }
};

type TaskType = {
  id: number;
  text: string;
  done: boolean;
};

const AddTask = () => {
  const [text, setText] = useState("");
  const dispatch = useContext(TaskDispatchContext);

  return (
    <>
      <input
        placeholder="Add task"
        value={text}
        onChange={(e) => setText(e.target.value)}
      />
      <button
        onClick={() => {
          setText("");
          dispatch?.({
            type: TaskActionType.ADD_TASK,
            payload: {
              id: nextId++,
              text,
              done: false,
            },
          });
        }}
      >
        Add
      </button>
    </>
  );
};

const Task = ({ task }: { task: TaskType }) => {
  const [isEditing, setIsEditing] = useState(false);
  const dispatch = useTaskDispatch();

  let taskContent;
  if (isEditing) {
    taskContent = (
      <>
        <input
          value={task.text}
          onChange={(e) =>
            dispatch?.({
              type: TaskActionType.CHANGE_TASK,
              payload: {
                ...task,
                text: e.target.value,
              },
            })
          }
        />
        <button onClick={() => setIsEditing(false)}>Save</button>
      </>
    );
  } else {
    taskContent = (
      <>
        {task.text}
        <button onClick={() => setIsEditing(true)}>Edit</button>
      </>
    );
  }

  return (
    <label>
      <input
        type="checkbox"
        checked={task.done}
        onChange={(e) =>
          dispatch?.({
            type: TaskActionType.CHANGE_TASK,
            payload: {
              ...task,
              done: e.target.checked,
            },
          })
        }
      />
      {taskContent}
      <button
        onClick={() =>
          dispatch?.({
            type: TaskActionType.DELETE_TASK,
            payload: {
              id: task.id,
            },
          })
        }
      >
        Delete
      </button>
    </label>
  );
};

const TaskList = () => {
  const tasks = useTasks();

  return (
    <ul>
      {tasks.map((task) => (
        <li key={task.id}>
          <Task task={task} />
        </li>
      ))}
    </ul>
  );
};

const initialTasks: TaskType[] = [
  { id: 0, text: "Visit Kafka Museum", done: true },
  { id: 1, text: "Watch a puppet show", done: false },
  { id: 2, text: "Lennon Wall pic", done: false },
];

let nextId = 3;

export default function Sample() {
  return (
    <>
      <TaskProvider>
        <h1>Prague itinerary</h1>
        <AddTask />
        <TaskList />
      </TaskProvider>
    </>
  );
}
