export 	interface HelpRequest {
	id: string;
	user_id: string;
	title: string;
	description: string;
	status: string;
	cargo_id: string;
}

export const Statuses = ['OPEN', 'IN PROGRESS', 'CLOSED'];



export function getCookie(name: string) {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${name}=`);
	if (parts.length === 2) return parts.pop()?.split(';').shift();
}