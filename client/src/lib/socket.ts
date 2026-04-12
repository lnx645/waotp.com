import { reactive } from "vue";
import { io } from "socket.io-client";

export const state = reactive<{
    connected : boolean,
    fooEvents : any[],
    barEvents : any[]
}>({
  connected: false,
  fooEvents: [],
  barEvents: []
});


export const socket = io("http://localhost:8080?token=dnine879hwq87eqasd7d9ah87ew",{
  transports : ["websocket"],
  path : "/realtime/ws"
});

socket.on("connect", () => {
  state.connected = true;
});

socket.on("disconnect", () => {
  state.connected = false;
});

socket.on("foo", (...args) => {
  return state.fooEvents.push(args);
});

socket.on("bar", (...args) => {
  return state.barEvents.push(args);
});
