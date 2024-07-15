<script lang="ts">
	import { Tabs, TabsList, TabsTrigger, TabsContent } from '$lib/components/ui/tabs';
	import {
		Card,
		CardContent,
		CardDescription,
		CardFooter,
		CardHeader,
		CardTitle
	} from '$lib/components/ui/card';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import { Textarea } from '$lib/components/ui/textarea';
	import { Button } from '$lib/components/ui/button';
	import { Calendar } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';
	import { goto } from '$app/navigation';

	let username = '';
	let password = '';
	let dateOfBirth = '';
	let full_name = '';
	let phone = '';
	let address = '';

	async function handleUserRegistration() {
		const response = await fetch('http://localhost:5000/api/v1/register', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				username,
				password,
				dateOfBirth,
				full_name,
				phone,
				address
			})
		});

		if (response.ok) {
			console.log('User registered successfully');
			toast.success('User registered successfully');

			goto('/');
		} else {
			console.error('Failed to register user');
			toast.error('Failed to register user', {
				description: response.status + ' ' + response.statusText
			});
		}
		console.log('User registration:', {
			username,
			password,
			dateOfBirth,
			fullName: full_name,
			phone,
			address
		});
	}
</script>

<div class="container mx-auto flex justify-center p-4">
	<Card>
		<CardHeader>
			<CardTitle>User Registration</CardTitle>
			<CardDescription>Create a new user account</CardDescription>
		</CardHeader>
		<CardContent>
			<form on:submit|preventDefault={handleUserRegistration} class="space-y-4">
				<div class="grid grid-cols-2 gap-4">
					<div class="space-y-2">
						<Label for="username">Username</Label>
						<Input id="username" bind:value={username} required />
					</div>
					<div class="space-y-2">
						<Label for="password">Password</Label>
						<Input id="password" type="password" bind:value={password} required />
					</div>
				</div>
				<div class="space-y-2">
					<Label for="fullName">Full Name</Label>
					<Input id="fullName" bind:value={full_name} required />
				</div>
				<div class="grid grid-cols-2 gap-4">
					<div class="space-y-2">
						<Label for="dateOfBirth">Date of Birth</Label>
						<div class="relative">
							<Input id="dateOfBirth" type="date" bind:value={dateOfBirth} required />
							<Calendar
								class="absolute right-3 top-1/2 -translate-y-1/2 transform text-gray-400"
								size={18}
							/>
						</div>
					</div>
					<div class="space-y-2">
						<Label for="phone">Phone</Label>
						<Input id="phone" type="tel" bind:value={phone} required />
					</div>
				</div>
				<div class="space-y-2">
					<Label for="address">Address</Label>
					<Textarea id="address" bind:value={address} required />
				</div>
				<Button
					type="submit"
					class="w-full"
					on:click={() => {
						handleUserRegistration;
					}}>Register User</Button
				>
			</form>
		</CardContent>
	</Card>
</div>
