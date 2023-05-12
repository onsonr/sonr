import React from 'react';

const PagePreview = (props) => {
	const {displayed} = props.document;
	if (!displayed?.slug?.current) {
		return <div>The page needs a slug before it can be previewed.</div>;
	}

	const slug = displayed?.slug?.current === 'frontpage' ? '' : displayed?.slug?.current;

	const url = new URL('/api/preview', location.origin);
		url.searchParams.set('slug', slug);
		url.searchParams.set('type', 'page');

	return (
		<div style={{width: '100%', height: '100%', position: 'relative', zIndex: 1}}>
			<div>
				<iframe
					src={url.toString()}
					style={{border: 0, height: '100%', left: 0, position: 'absolute', top: 0, width: '100%'}}
				/>
			</div>
		</div>
	);
};

export default PagePreview;
