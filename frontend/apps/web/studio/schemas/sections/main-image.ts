import {RiImageLine} from 'react-icons/ri';
import {defineField} from 'sanity';

const mainImage = defineField({
	name: 'mainImage',
	title: 'Image',
	description: 'Image',
	type: 'image',
	icon: RiImageLine,
	options: {
		hotspot: true
	},
	fields: [
		{
			name: 'alt',
			title: 'Alternative text',
			type: 'string',
			description: 'Important for SEO and accessiblity.',
			validation: (Rule) => Rule.error('You have to fill out the alternative text.').required()
		}
	],
	preview: {
		select: {
			title: 'alt',
			media: 'asset'
		},
		prepare({media, title}) {
			return {
				media,
				title
			};
		}
	}
});

export default mainImage;
