<script lang="ts">
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import { toast } from 'svelte-sonner';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { Package, LogIn, LogOut, CircleUserRound, SquareUserRound } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { getCookie } from '../types/types';

	let userRole: string = 'guest';
	let userId: number = 0;
	let username: string = '';
	let password: string = '';

	$: isLoggedIn = userRole !== 'guest';
	$: showAdminPanel = userRole === '3' || userRole === '4';

	onMount(getUserDataFromCookies);

	function getUserDataFromCookies() {
		const savedToken = getCookie('authToken');
		const savedRole = getCookie('userRole');
		if (savedToken && savedRole) {
			userRole = savedRole;
		}
	}

	async function handleLogin() {
		try {
			const response = await fetch('http://localhost:5000/api/v1/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, password })
			});

			if (response.ok) {
				const data = await response.json();
				setCookie('authToken', data.token, 7);
				toast.success('Login successful', {
					description: 'Вітаємо!'
				});
			} else {
				console.error('Login failed');
				toast.error('Login failed', {
					description: response.status + ' ' + response.statusText
				});
			}

			const identityResponse = await fetch('http://localhost:5000/api/v1/identity', {
				headers: {
					Authorization: `Bearer ${getCookie('authToken')}`
				}
			});

			if (identityResponse.ok) {
				const data = await identityResponse.json();
				userRole = data.role;
				userId = data.user_id;
				setCookie('userRole', userRole, 7);
				setCookie('userId', userId.toString(), 7);
				getUserDataFromCookies();
			} else {
				console.error('Failed to get user identity');
				toast.error('Identification failed', {
					description: response.status + ' ' + response.statusText
				});
			}
		} catch (error) {
			console.error('Error during login:', error);
			toast.error('Error');
		}
	}

	function handleLogout() {
		deleteCookie('authToken');
		deleteCookie('userRole');
		userRole = 'guest';
	}

	function setCookie(name: string, value: string, days: number) {
		const expires = new Date(Date.now() + days * 864e5).toUTCString();
		document.cookie = `${name}=${encodeURIComponent(value)}; expires=${expires}; path=/`;
	}

	function deleteCookie(name: string) {
		document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
	}
</script>

<header class="bg-primary text-primary-foreground shadow-lg">
	<div class="container mx-auto flex items-center justify-between px-4 py-2">
		<div class="flex items-center space-x-2">
			<Package class="h-8 w-8" />
			<span class="text-2xl font-bold">AidTracker</span>
		</div>
		<nav class="flex items-center">
			<ul class="mr-4 flex space-x-4">
				<li><Button variant="ghost" class="text-lg" on:click={() => goto('/')}>Головна</Button></li>
				<li>
					<Button variant="ghost" class="text-lg" on:click={() => goto('/aid-requests')}
						>Заявки на допомогу</Button
					>
				</li>
				<li><Button variant="ghost" class="text-lg">Про нас</Button></li>
				{#if showAdminPanel}
					<li>
						<Button variant="secondary" class="text-lg" on:click={() => goto('/control-panel')}
							>Панель керування</Button
						>
					</li>
				{/if}
			</ul>
		</nav>
		<div>
			{#if isLoggedIn}
				<DropdownMenu.Root>
					<DropdownMenu.Trigger>
						<Button variant="secondary" class="flex items-center text-lg">
							<CircleUserRound class="mr-2" />
							Акаунт
						</Button>
					</DropdownMenu.Trigger>
					<DropdownMenu.Content>
						<DropdownMenu.Group>
							<DropdownMenu.Item on:click={() => goto('/profile')} class="flex items-center">
								<SquareUserRound />&nbsp; Профіль
							</DropdownMenu.Item>
							<DropdownMenu.Item class="flex items-center" on:click={handleLogout}>
								<LogOut />&nbsp; Вийти
							</DropdownMenu.Item>
						</DropdownMenu.Group>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
			{:else}
				<Dialog.Root>
					<Dialog.Trigger class={buttonVariants({ variant: 'secondary' })}>
						<div class="flex items-center text-lg">
							<LogIn class="mr-2 h-4 w-4" />
							Увійти
						</div>
					</Dialog.Trigger>
					<Dialog.Content class="sm:max-w-[425px]">
						<Dialog.Header>
							<Dialog.Title class="text-2xl">Увійдіть у свій акаунт</Dialog.Title>
						</Dialog.Header>
						<div class="grid gap-4 py-4">
							<div class="grid grid-cols-4 items-center gap-4">
								<Label for="email" class="text-right">Email</Label>
								<Input
									id="email"
									type="email"
									placeholder="email@example.com"
									class="col-span-3"
									bind:value={username}
								/>
							</div>
							<div class="grid grid-cols-4 items-center gap-4">
								<Label for="password" class="text-right">Пароль</Label>
								<Input
									type="password"
									id="password"
									placeholder="Your password"
									class="col-span-3"
									bind:value={password}
								/>
							</div>
						</div>
						<Dialog.Footer>
							<Button type="submit" on:click={handleLogin}>Увійти</Button>
						</Dialog.Footer>
					</Dialog.Content>
				</Dialog.Root>
			{/if}
		</div>
	</div>
</header>
