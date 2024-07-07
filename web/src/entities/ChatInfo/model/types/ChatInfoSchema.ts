export interface ChatInfo {
    id: number;
    name: string;
    last_message: string;
    last_message_time: string;
    unread: number;
}

export interface ChatListSchema {
    chats: ChatInfo[];
    isLoading: boolean;
    error?: string;
}
