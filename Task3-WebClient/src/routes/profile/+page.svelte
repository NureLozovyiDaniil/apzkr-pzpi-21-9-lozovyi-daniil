<script lang="ts">
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';
	import * as AlertDialog from '$lib/components/ui/alert-dialog';
	import { Tabs, TabsList, TabsTrigger, TabsContent } from '$lib/components/ui/tabs';
	import { Input } from '$lib/components/ui/input';
	import { Textarea } from '$lib/components/ui/textarea';
	import { Label } from '$lib/components/ui/label';
	import { User, Shield, UserRoundCog } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';
	import { Badge } from '$lib/components/ui/badge/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { getCookie } from '../../types/types';

	let userRole: string = getCookie('userRole') as '1' | '2' | '3' | '4';

	onMount(getUser);

	async function getUser() {
		const response = await fetch('http://localhost:5000/api/v1/user', {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${getCookie('authToken')}`
			}
		});

		if (response.ok) {
			const data = await response.json();
			console.log(data);
			userRole = getCookie('userRole') as '1' | '2' | '3' | '4';
			profileForm = {
				username: data.username,
				fullName: data.full_name,
				address: data.address,
				phoneNumber: data.phone,
				role: userRole
			};
		} else {
			console.error('Failed to get user data');
			toast.error('Failed to get user data', {
				description: response.status + ' ' + response.statusText
			});
		}
	}

	let activeSection: 'profile' | 'security' | 'changeRole' = 'profile';

	let profileForm = {
		username: '',
		fullName: '',
		address: '',
		phoneNumber: '',
		role: ''
	};

	let securityForm = {
		currentPassword: '',
		newPassword: '',
		confirmPassword: ''
	};

	function validateSecurityForm() {
		// Add your validation logic here
		console.log('Validating security form');
	}

	async function handleVolunteer() {
		const response = await fetch('http://localhost:5000/api/v1/volunteer', {
			method: 'POST',
			headers: {
				Authorization: `Bearer ${getCookie('authToken')}`
			}
		});

		if (response.ok) {
			console.log('User role changed to Volunteer');
			toast.success('User role changed to Volunteer');
		} else {
			console.error('Failed to change user role');
			toast.error('Failed to change user role', {
				description: response.status + ' ' + response.statusText
			});
		}
	}

	let orgName = '';
	let orgDescription = '';

	async function handleOrg() {
		console.log('Handling organization form');
		const response = await fetch('http://localhost:5000/api/v1/organization', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${getCookie('authToken')}`
			},
			body: JSON.stringify({
				name: orgName,
				description: orgDescription
			})
		});

		if (response.ok) {
			console.log('User role changed to Organization');
			toast.success('User role changed to Organization');
		} else {
			console.error('Failed to change user role');
			toast.error('Failed to change user role', {
				description: response.status + ' ' + response.statusText
			});
		}
	}
</script>

<div class="container mx-auto p-4">
	<h1 class="mb-4 text-2xl font-bold">Settings</h1>
	<p class="mb-6 text-gray-600">Manage your account settings and set e-mail preferences.</p>

	<div class="flex">
		<!-- Left sidebar -->
		<div class="w-1/4 pr-4">
			<Button
				variant={activeSection === 'profile' ? 'default' : 'ghost'}
				class="mb-2 w-full justify-start"
				on:click={() => (activeSection = 'profile')}
			>
				<User class="mr-2 h-4 w-4" />
				Profile
			</Button>
			<Button
				variant={activeSection === 'security' ? 'default' : 'ghost'}
				class="w-full justify-start"
				on:click={() => (activeSection = 'security')}
			>
				<Shield class="mr-2 h-4 w-4" />
				Security
			</Button>
			<Separator class="my-4" />
			{#if userRole != '4' && userRole != '3'}
				<Button
					variant={activeSection === 'changeRole' ? 'default' : 'ghost'}
					class="w-full justify-start"
					on:click={() => (activeSection = 'changeRole')}
				>
					<UserRoundCog class="mr-2 h-4 w-4" />
					Change role
				</Button>
			{/if}
		</div>

		<!-- Main content -->
		<div class="w-3/4">
			{#if activeSection === 'profile'}
				<h2 class="mb-4 text-xl font-semibold">Profile</h2>
				<Badge class="text-md">
					{#if profileForm.role === '1'}
						User
					{:else if profileForm.role === '2'}
						Volunteer
					{:else if profileForm.role === '3'}
						Organization
					{:else if profileForm.role === '4'}
						Admin
					{/if}
				</Badge>
				<form class="space-y-4">
					<div>
						<Label for="username">Username</Label>
						<Input id="username" bind:value={profileForm.username} />
					</div>
					<div>
						<Label for="fullName">Full Name</Label>
						<Input id="fullName" bind:value={profileForm.fullName} />
					</div>
					<div>
						<Label for="address">Address</Label>
						<Textarea id="address" bind:value={profileForm.address} />
					</div>
					<div>
						<Label for="phoneNumber">Phone Number</Label>
						<Input id="phoneNumber" bind:value={profileForm.phoneNumber} />
					</div>
					<Button type="submit">Save Changes</Button>
				</form>
			{:else if activeSection === 'security'}
				<h2 class="mb-4 text-xl font-semibold">Security</h2>
				<form class="space-y-4" on:submit|preventDefault={validateSecurityForm}>
					<div>
						<Label for="currentPassword">Current Password</Label>
						<Input id="currentPassword" type="password" bind:value={securityForm.currentPassword} />
					</div>
					<div>
						<Label for="newPassword">New Password</Label>
						<Input id="newPassword" type="password" bind:value={securityForm.newPassword} />
					</div>
					<div>
						<Label for="confirmPassword">Confirm New Password</Label>
						<Input id="confirmPassword" type="password" bind:value={securityForm.confirmPassword} />
					</div>
					<Button type="submit">Change Password</Button>
				</form>
			{:else if activeSection === 'changeRole'}
				<h2 class="mb-4 text-xl font-semibold">Change Role</h2>
				<Tabs class="w-full">
					<TabsList class="grid w-full grid-cols-2">
						{#if userRole === '1'}
							<TabsTrigger value="volunteer">Become volunteer</TabsTrigger>
						{/if}
						<TabsTrigger value="organization">Become organization</TabsTrigger>
					</TabsList>
					{#if userRole === '1'}
						<TabsContent value="volunteer">
							<h2 class="mb-4 text-xl font-semibold">Volunteer</h2>
							<p class="mb-4 text-gray-600">
								By becoming a volunteer, you will be able to help people in need and make a
								difference in the world.
							</p>
							<AlertDialog.Root>
								<AlertDialog.Trigger>
									<Button variant="secondary">Become Volunteer</Button>
								</AlertDialog.Trigger>
								<AlertDialog.Content>
									<AlertDialog.Header>
										<AlertDialog.Title>Ви точно впевнені?</AlertDialog.Title>
										<AlertDialog.Description>
											Ця дія необратима. Після зміни ващої ролі ви не зможете повернутися назад.
										</AlertDialog.Description>
									</AlertDialog.Header>
									<AlertDialog.Footer>
										<AlertDialog.Cancel>Ні, дякую</AlertDialog.Cancel>
										<AlertDialog.Action on:click={handleVolunteer}>Продовжити</AlertDialog.Action>
									</AlertDialog.Footer>
								</AlertDialog.Content>
							</AlertDialog.Root>
						</TabsContent>
					{/if}
					<TabsContent value="organization">
						<h3 class="mb-4 text-xl font-semibold">Organization</h3>
						<div class="px">
							<p class="mb-4 text-gray-600">
								By becoming an organization, you will be able to create and manage aid requests and
								coordinate humanitarian efforts.
							</p>
							<div class="space-y-2">
								<Label for="orgName">Organization Name</Label>
								<Input id="orgName" bind:value={orgName} required />
							</div>
							<div class="mb-3 space-y-2">
								<Label for="orgDescription">Organization Description</Label>
								<Textarea id="orgDescription" bind:value={orgDescription} required />
							</div>
							<Button on:click={handleOrg}>Стати організацією</Button>
						</div>
					</TabsContent>
				</Tabs>
			{/if}
		</div>
	</div>
</div>
