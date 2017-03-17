/* @flow */

import React from 'react';
import WorkspaceLayout from 'pages/WorkspaceLayout';

class WorkspaceMaterialContainer extends React.Component {
  render() {
    return (
      <WorkspaceLayout
        isWorkspaceLoading
        activeWorkspaceTab="material"
      />
    );
  }
}

export default WorkspaceMaterialContainer;
