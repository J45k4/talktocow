import { ReactVar, ReactVars } from "./react-var";
import { Chatroom } from "./types";

export const chatroomsState = new ReactVar([])
export const creatingChatroomState = new ReactVar(false)
export const selectedUsersState = new ReactVar([])

export const myChatroomsState = new ReactVar<null | Chatroom[]>(null)
export const chatroomState = new ReactVars<string, Chatroom>()